<template lang="pug">
  .view
    GraphViz(:dotData='graph')
    button(type="button" @click='reloadGraph()')
      | Recargar
</template>

<script>
import { get } from '@/services/graph'
import GraphViz from '@/components/GraphViz.vue'
export default {
  name: 'view',
  data () {
    return {
      graph: 'digraph { a-> b }'
    }
  },
  components: {
    GraphViz
  },
  methods: {
    async reloadGraph () {
      try {
        const retreivedData = await get()
        this.graph = retreivedData.data
      } catch (e) {
        console.log(e)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
$view-max-width: 600px;

.graph {
  max-width: $view-max-width;
}
</style>
