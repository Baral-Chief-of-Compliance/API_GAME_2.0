<template>
    <v-container class="mx-10">
        <v-row>
            <v-col cols="4">
                <span class="title_for_inf">Имя: {{ props.heroName }}</span>
            </v-col>
            <v-col  cols="4">
                <span class="title_for_inf">Тип:{{ props.hetoType }}</span>
            </v-col>
            <v-col  cols="4">
                
            </v-col>
        </v-row>

        <v-table class="mt-15 text-center">
            <thead >
                <tr>
                    <th class="text-center">
                    <span class="title_for_inf">Доверие</span>
                    </th>
                    <th class="text-center">
                        <span class="title_for_inf">Фрагмент истории</span>
                    </th>
                </tr>
            </thead>

            <tbody >
                <tr
                    v-for="stor in stories.stories"
                    :key="stor.id_stgp"
                >
                    <td>{{ stor.confidence_gp }}</td>
                    <td>{{ stor.story }}</td>
                </tr>
            </tbody>
        </v-table>

        <button-add title="Добавить фрагмент истории" />
        

    </v-container>
</template>


<script setup>
import axios from 'axios';

    const url_api = inject('url_api')
    const stories = reactive({
        stories: []
    })

    const props = defineProps(
        {
            heroName: String,
            hetoType: String,
            heroId: Number
        }
    )

    function getHeroInf(id){
        axios.get(`${url_api}game_persons/${id}/information`)
        .then(
            function(response){
                stories.stories = response.data.stories
            }
        )
    }

    getHeroInf(props.heroId)
</script>

<style>
    .title_for_inf{
        font-size: 24px;
        font-weight: bold;
        color: black
    }
</style>