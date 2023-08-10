<template>
    <v-container>


        <TitlePage title="Создание сюжет" />
        
        <ButtonBack title="назад к сюжету" @click="BackStories" />

        <v-container>
            <TextField label="Назвние главы" @customChange="getStoryName" />
            <TextArea label="Содержание" @customChange="getStoryContent" />
            <ButtonComponents class="mt-5" label="Добавить сюжет" @click="createStory" />
        </v-container>
    </v-container>
</template>

<script setup>
import axios from 'axios';


    const Story = reactive({
        name: "",
        description: ""
    })
    const url_api = inject("url_api")
    const router = useRouter()

    function getStoryName(event){
        Story.name = event
    }

    function getStoryContent(event){
        Story.description = event 
    }

    function createStory(){
        axios.post(`${url_api}stories`, {
            name_st: Story.name,
            content_st: Story.description
        }).then(
            function(){
                router.push({
                    name: 'stories'
                })
            }
        )
    }

    function BackStories(){
        router.push({name: 'stories'})
    }

</script>