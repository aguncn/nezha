<template>
  <environment-form :id=parseInt(this.$route.params.id) :readonly="false" :opType="'update'" />
</template>

<script>
import EnvironmentForm from './components/form'

export default {
  name: 'EditProject',
  components: { EnvironmentForm }
}
</script>
