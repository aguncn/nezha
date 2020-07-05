package deploy

import (
	"strconv"
	"time"

	"github.com/aguncn/nezha/common/datasource"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	baseRep "github.com/aguncn/nezha/repository"
	"github.com/jinzhu/gorm"
)

type DeployRepository struct {
	Log    logger.ILogger         `inject:""`
	Base   baseRep.BaseRepository `inject:"inline"`
	Source datasource.IDb         `inject:""`
}

func (deploy *DeployRepository) GetDeploy(where interface{}) *models.Deploy {
	var modDeploy models.Deploy
	db := deploy.Source.DB().Where(where)
	if err := db.First(&modDeploy).Error; err != nil {
		deploy.Log.Errorf("未找到相关Deploy", err)
	}
	return &modDeploy
}

func (deploy *DeployRepository) AddDeploy(modDeploy *models.Deploy) bool {
	if err := deploy.Base.Create(modDeploy); err != nil {
		deploy.Log.Errorf("添加文章失败", err)
		return false
	}
	return true
}

func (deploy *DeployRepository) UpdateDeploy(modDeploy *models.Deploy) bool {
	if err := deploy.Base.Save(modDeploy); err != nil {
		deploy.Log.Errorf("更新应用失败", err)
		return false
	}
	return true
}

func (deploy *DeployRepository) UpdateDeployByCol(modDeploy *models.Deploy, where map[string]interface{}) bool {

	if err := deploy.Source.DB().Model(&modDeploy).Updates(where).Error; err != nil {
		deploy.Log.Errorf("更新应用失败", err)
		return false
	}
	return true
}

func (deploy *DeployRepository) GetDeploys(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.Deploy {

	var deploys []models.Deploy
	db := deploy.Source.DB().Model(&deploys).Where(&deploys)
	db = db.Preload("Environment", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	}).Preload("K8s", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	})

	db = db.Order("updated_at desc")
	db = db.Where(where)
	err := db.Offset((PageNum - 1) * PageSize).Limit(PageSize).Find(&deploys).Error
	if err != nil {
		deploy.Log.Errorf("获取Applications失败", err)
	}
	err = db.Count(total).Error
	if err != nil {
		deploy.Log.Errorf("查询总数出错", err)
		// return err
	}
	return &deploys
}

func (deploy *DeployRepository) GetOkErrorCount(where string) interface{} {
	type OkError struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
	var okCount int
	var errorCount int
	var okError OkError
	var okErrors []OkError

	var deploys []models.Deploy
	db1 := deploy.Source.DB().Model(&deploys).Where(&deploys)
	db1 = db1.Order("updated_at desc")
	db1 = db1.Where(where).Where("result LIKE '%success%'")

	err := db1.Count(&okCount).Error
	if err != nil {
		deploy.Log.Errorf("查询总数出错", err)
		// return err
	}
	okError.Name = "成功"
	okError.Value = okCount
	okErrors = append(okErrors, okError)

	db2 := deploy.Source.DB().Model(&deploys).Where(&deploys)
	db2 = db2.Order("updated_at desc")
	db2 = db2.Where(where).Where("result LIKE '%error%'")

	err = db2.Count(&errorCount).Error
	if err != nil {
		deploy.Log.Errorf("查询总数出错", err)
		// return err
	}
	okError.Name = "失败"
	okError.Value = errorCount
	okErrors = append(okErrors, okError)
	return okErrors
}

func (deploy *DeployRepository) GetBetweenCount(where string, cycle, count int) interface{} {
	// cycle表示取值周期，比如，以周，月，年为取值周期
	// count表示取值个数，比如，取4周，4月，4年
	var cycleBetween string
	var countBetween int
	type BetweenCount struct {
		Cycle []string `json:"cycle"`
		Count []int    `json:"count"`
	}
	var betweenCount BetweenCount

	var deploys []models.Deploy
	for i := count; i >= 0; i-- {
		db := deploy.Source.DB().Model(&deploys).Where(&deploys)
		now := time.Now()

		cycleStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, 0-(i+1)*cycle)
		cycleEndDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, 0-(i)*cycle)

		db = db.Order("updated_at desc")
		db = db.Where(where).Where("created_at BETWEEN ? AND ?", cycleStartDate, cycleEndDate)

		err := db.Count(&countBetween).Error
		if err != nil {
			deploy.Log.Errorf("查询总数出错", err)
			// return err
		}
		cycleBetween = strconv.Itoa(i) + "周前"
		if i == 0 {
			cycleBetween = "本周"
		}

		betweenCount.Cycle = append(betweenCount.Cycle, cycleBetween)
		betweenCount.Count = append(betweenCount.Count, countBetween)
	}

	return betweenCount
}
func (deploy *DeployRepository) ExistDeployByName(where interface{}) bool {
	var modDeploy models.Deploy
	var whereDeploy = models.Deploy{Name: where.(*models.Deploy).Name}
	other := map[string]string{}
	err := deploy.Base.First(whereDeploy, &modDeploy, other)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		deploy.Log.Error(err)
		return false
	}
	return true
}

func (deploy *DeployRepository) DeleteDeploy(id uint) bool {
	var modDeploy models.Deploy
	if err := deploy.Base.DeleteByID(modDeploy, id); err != nil {
		deploy.Log.Errorf("删除应用失败", err)
		return false
	}
	return true
}
