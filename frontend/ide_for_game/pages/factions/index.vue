<template>
    <div class="FactionsMain">
        <div class="factionsPage">
        <MainPage  
            title="Фракции" 
            :data="factions" 
            title_for_batton="Добавить фракции"
            path="factions"
        />
    </div>
    </div>
</template>

<script>
import axios from 'axios';

    export default{
        data: () => ({
            factions: []
        }),

        methods: {
            get_all_factions(){
                axios.get("http://127.0.0.1:5000/the_stone_sword/api/v1.0/faction")
                .then(
                    response => (
                        this.factions = response.data.factions
                    )
                )
            }
        },

        mounted(){
            this.get_all_factions()
        }
    }
</script>