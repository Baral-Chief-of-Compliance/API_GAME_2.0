<template>
    <v-container>
        <TitlePage :title="makeTitle(gamePerson.name)" />
        <ButtonBack title="назад к персонажам" @click="BackPersons" />
        <Tubs 
            :tab="tab"
            :tabs="['Информация', 'Характеристики','Иконка', 'Анимации']" 
            @get-index="updateIndex"
        />

        <GamePersonInf 
            :hero-name="gamePerson.name" 
            :heto-type="gamePerson.type"
            :hero-id="router.params.id"
            v-if="tab == 0" 
        />
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
    const Router = useRouter()


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

    function BackPersons(){
            Router.push({name: "game_persons"})
    }


    onMounted(() => {
        getInfGamePerson(router.params.id)
    })

</script>