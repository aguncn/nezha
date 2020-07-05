<template>
  <environment-form :id=parseInt(this.$route.params.id) :readonly="false" :opType="'create'" />
</template>

<script>
import EnvironmentForm from './components/form'

export default {
  name: 'CreateEnvironment',
  components: { EnvironmentForm }

}
</script>
