<template>
    <v-container>
        <TitlePage :title="makeTitle(gamePerson.name)" />
        <Tubs 
            :tab="tab"
            :tabs="['Информация', 'Характеристики','Иконка', 'Анимации']" 
            @get-index="updateIndex"
        />

        <GamePersonInf v-if="tab == 0" />
        <GamePersonCharacteristics v-else-if="tab == 1" />
        <GamePersonGamePesonIcon v-else-if="tab == 2"/>
        <GamePersonGamePersonsAnimation v-else-if="tab == 3"/>

    </v-container>

   
    
</template>

<script setup>
import axios from 'axios'; 

    const tab = ref("0")
    const url_api = inject('url_api')
    const gamePerson = reactive({
        name: "",
        type: ""
    })
    const router = useRoute()


    function updateIndex(index){
        tab.value = index
    }

    function makeTitle(name){
        return `Игровой персонаж: ${name}`
    }

    function getInfGamePerson(id_gp){
        axios.get(`${url_api}game_persons/${id_gp}`)
        .then(
            function(response){
                console.log(response)
                gamePerson.name = response.data.name_gp;
                gamePerson.type = response.data.type_gp;
            }
        )
    }


    onMounted(() => {
        getInfGamePerson(router.params.id)
    })

</script>