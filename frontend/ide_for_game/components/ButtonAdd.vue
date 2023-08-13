<template>
            <v-btn 
                    v-model="dialog.status"
                    color="indigo"
                    block
                    class="mt-15"
                    v-bind="props"
                >
                    {{ title }}

                <v-dialog
                    v-model="dialog"
                    activator="parent"
                    width="auto"
                >
                    <GamePersonCardsGamePersonsAddInf :close="closeDialog" :dialog="dialog" />
                </v-dialog>
            </v-btn> 
</template>

<script setup>

    const props = defineProps(
        {
            title: String
        }
    )

</script>


<script>
  export default {
    data () {
      return {
        dialog: false,
      }
    },

    methods: {
        closeDialog(){
            this.dialog = false
        } 
    }
  }
</script>
