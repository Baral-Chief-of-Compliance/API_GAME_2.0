<template>
    <v-container>
        <TitlePage title="Создание персонажа" />

        <ButtonBack title="назад к персонажам" @click="BackPersons" />

        <v-container>
            <TextField label="Имя" @customChange="getGamePersonName" />
            <Select 
                class="mt-5" 
                label="Тип" 
                @inputed="getGamePersonType"
                :items="[
                    'Игровой персонаж', 
                    'Рядовой враг',
                    'Мини-босс',
                    'Босс',
                    'NPS'
                ]" 
            />
            <ButtonComponents class="mt-5" label="Добавить персонажа" @click="createHero" />
        </v-container>
    </v-container>
</template>

<script setup>
import axios from 'axios';

        const typePerson = reactive({ name: ""})
        const url_api = inject('url_api')
        const gamePersonName = reactive({name: ""})
        const router = useRouter()

        function getGamePersonType(event){
            typePerson.name = event;
        }

        function getGamePersonName (event){
            gamePersonName.name = event
        }

        function createHero(){
            axios.post(`${url_api}game_persons`, {
                name_gp: gamePersonName.name,
                type_gp: typePerson.name
            }).then(
                function(){
                    router.push({
                        name: 'game_persons'
                    })
                }
            )
        }

        function BackPersons(){
            router.push({name: "game_persons"})
        }
</script>