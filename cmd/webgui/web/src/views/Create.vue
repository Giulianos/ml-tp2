<template lang="pug">
.create-view
  form.build(@submit.prevent='sendCSV')
    textarea(v-model='csvData' placeholder='Insertar el dataset en CSV aca')
    button
      | Crear árbol
  GraphViz(dotData='digraph { a-> b }').graph
</template>

<script>
// @ is an alias to /src
import GraphViz from '@/components/GraphViz.vue'
import { create } from '@/services/tree'

export default {
  data () {
    return {
      csvData: '',
      predAttr: 'disfruta'
    }
  },
  methods: {
    sendCSV () {
      create(this.csvData, this.predAttr)
    }
  },
  name: 'home',
  components: {
    GraphViz
  }
}
</script>

<style lang="scss" scoped>
$view-max-width: 600px;

.create-view {
  display: flex;
  align-items: center;
  flex-direction: column;
}

.build {
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: $view-max-width;

  textarea {
    height: 100px;
  }
}

.graph {
  max-width: $view-max-width;
}
</style>
