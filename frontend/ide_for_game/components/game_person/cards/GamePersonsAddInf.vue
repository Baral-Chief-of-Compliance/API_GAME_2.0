<template>
    <v-card>
        <v-card-title>
            Добавить Фрагменты истории 
        </v-card-title>

        <v-card-subtitle>
            Добавьте фрагмент истории и очки доверия, при котороых этот фргамент будет доступен игроку      
        </v-card-subtitle>

        <v-card-text>
            <NumberField label="Доверие" @customChange="getStoryConfidence" />
            <TextArea label="Фрагмент истории" @customChange="getStory" />
        </v-card-text>

        <v-card-actions class="mx-15">
                <v-btn variant="outlined" color="indigo" @click="sendHeroStory" >
                    Добавить фрагмент истории
                </v-btn>
                <v-spacer></v-spacer>
                <v-btn variant="outlined"  color="red" @click="props.close">
                    закрыть
                </v-btn>
        </v-card-actions>
    </v-card>
</template>

<script setup>
import axios from 'axios';

    const props = defineProps({
        dialog: Boolean,
        close: Function
    })

    const route = useRoute()

    const HeroStory = reactive({
        confidence: 0,
        story: ""
    })

    const url_api = inject("url_api")

    function getStory(event){
        HeroStory.story = event
    }

    function getStoryConfidence(event){
        HeroStory.confidence = event
    }

    function sendHeroStory(){
        console.log(route.params.id)

        props.close()
    }
</script>