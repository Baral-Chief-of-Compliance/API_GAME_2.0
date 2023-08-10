<template>
    <div class="StoriesPage">
        
        <MainPage 
            title="Сюжет" 
            title_for_batton="Добавить сюжет"
            path="stories/create"
            :data="stories"
        />
    </div>
</template>

<script setup>
import axios from 'axios';


const stories = ref([])
const url_api = inject("url_api")


function getStories(){
    axios.get(`${url_api}stories`)
    .then(
        function(response){
            stories.value = response.data.stories
        }
    )
}

getStories()


</script>