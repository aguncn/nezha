<template>
  <div class="page-container">
	  <el-tabs type="card" v-model="activeName" @tab-click="handleTabClick">
		  <el-tab-pane label="发布单" name="deploy" >
			  <deploy-list :id=parseInt(this.$route.params.id) v-if="tabRefresh.deploy"/>
		  </el-tab-pane>
	    <el-tab-pane label="历史部署" name="history">
			  <history-list :id=parseInt(this.$route.params.id) v-if="tabRefresh.history"/>
		  </el-tab-pane>
		  <el-tab-pane label="线上Pod" name="pod">
			  <pod-list :id=parseInt(this.$route.params.id) v-if="tabRefresh.pod"/>
		  </el-tab-pane>
	  </el-tabs>
  </div>
</template>

<script>
	import DeployList from "@/views/deploy/deploylist"
	import HistoryList from "@/views/deploy/historylist"
	import PodList from "@/views/deploy/podlist"
	import { format } from "@/utils/datetime"
	export default {
    name: 'DeployIndex',
		components:{
			PodList,
			DeployList,
			HistoryList
		},
		data() {
			return {
				activeName: "deploy",
				tabRefresh: {
					deploy: true,
					history: false,
					pod: false
				}
			}
		},
		methods: {
			handleTabClick: function(tab, event){
				switch (this.activeName) {
					case 'deploy':
						this.switchTab('deploy')
						console.log(this.activeName);
						break;
					case 'history':
						this.switchTab('history')
						console.log(this.activeName);
						break;
					case 'pod':
						this.switchTab('pod')
						console.log(this.activeName);
						break;
					default:
						console.log('wrong choice');

				}
			},
			switchTab: function(tab) {
				for (let [key, value] of Object.entries(this.tabRefresh)) {
					if (key == tab) {
						this.tabRefresh[key] = true
					} else {
						this.tabRefresh[key] = false
					}
				}
			}
		}
	}
</script>
<style scoped>
  .page-container {
    text-align: left;
    font-size: 16px;
    color: rgb(20, 89, 121);
  }
</style>
