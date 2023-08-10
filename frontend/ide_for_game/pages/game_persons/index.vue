<template>
    <div class="gamePersonsPage">
        <MainPage  
            title="Игровые персонажи" 
            :data="game_persons" 
            title_for_batton="Добавить персонажа"
            path="game_persons"
        />
    </div>
</template>

<script>
import axios from 'axios';

    export default{
        data: ()=> ({
            game_persons: []
        }),

        methods: {
            get_all_persons(){
                axios.get("http://127.0.0.1:5000/the_stone_sword/api/v1.0/game_persons")
                .then(
                    response => (
                        this.game_persons = response.data.game_persons
                    )
                )
            }
        },

        mounted(){
            this.get_all_persons()
        }
    }
</script>