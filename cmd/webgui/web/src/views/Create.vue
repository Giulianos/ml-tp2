<template lang="pug">
.create-view
  form.build(@submit.prevent='sendCSV')
    textarea.w100.mb-2(v-model='csvData' placeholder='Conjunto de datos en CSV')
    select.w100.mb-1(v-model='predAttr' placeholder='Atributo a predecir')
      option(v-for='attr in attributes' :value='attr')
        | {{attr}}
    input.w100.mb-1(v-model='minNodeCount' placeholder='Mín. cant. de ejemplos por nodo')
    select.w100.mb-2(v-model='gainFunc' placeholder="Función de ganancia")
      option(select="selected" value="shannon")
        | Entropía de Shannon
      option(value="gini")
        | Índice de Gini
    button.button.mb-1
      | Crear árbol
</template>

<script>
// @ is an alias to /src
import GraphViz from '@/components/GraphViz.vue'
import { create } from '@/services/tree'

export default {
  data () {
    return {
      csvData: '',
      predAttr: '',
      minNodeCount: '',
      gainFunc: 'shannon'
    }
  },
  methods: {
    sendCSV () {
      create(this.csvData, this.predAttr, this.gainFunc, this.minNodeCount)
    }
  },
  computed: {
    attributes() {
      return this.csvData.length ? this.csvData.split('\n')[0].split(',') : []
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
  align-items: center;

  textarea {
    height: 100px;
  }
}

.graph {
  max-width: $view-max-width;
}

.mb-1 {
  margin-bottom: 10px;
}

.mb-2 {
  margin-bottom: 20px;
}

.w100 {
  width: 100%;
}

.button {
  background-color: #42b983;
  color: white;
  height: 40px;
  border-radius: 20px;
  border: none;
  font-size: 16px;
  width: 250px;
}
</style>
