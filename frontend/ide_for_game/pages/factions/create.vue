<template>
   <v-container>
        <TitlePage title="Создание фракции" />

        <ButtonBack title="назад к фракциям" @click="BackFactions" />

        <v-container>
            <TextField label="Название фракции" @customChange="getFactionName" />
            <TextArea label="Описание фракции" @customChange="getFactionDescription" />
            <ButtonComponents class="mt-5" label="Добавить фракцию" @click="createFuction" />
        </v-container>
   </v-container> 
</template>

<script setup>
import axios from 'axios';


    const faction = reactive({
        name: "",
        description: ""
    })
    const url_api = inject('url_api')
    const router = useRouter()


    function getFactionName(event){
        faction.name = event;
    }

    function getFactionDescription(event){
        faction.description = event;
    }

    function createFuction(){
        axios.post(`${url_api}faction`, {
            name_f: faction.name,
            description_f: faction.description
        }).then(
            function(){
                router.push({
                    name: 'factions'
                })
            }
        )
    }

    function BackFactions(){
        router.push({name: "factions"})
    }
</script>