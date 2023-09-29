<template>
    <v-container>
        <TitlePage title="Задачи" />

        <v-container class="mx-15">

            <v-dialog width="700">
                <template v-slot:activator="{ props }">
                    <v-btn v-bind="props" color="green-darken-4" block>Добавить задачу</v-btn>
                </template>

                <template v-slot:default="{ isActive }">
                    
                    <v-card title="Добавить задачу">
                        <v-form @submit.prevent="submit">
                            <v-card-text>
                            
                                    <TextField :rules="this.ruleName" label="Название задачи" @customChange="getTaskName"/>
                                    <Select
                                        :rules="this.ruleArea"
                                        @inputed="getTaskArea"
                                        label="Область работы"
                                        :items="[
                                            'программирование',
                                            'рисование',
                                            'сюжетобразование',
                                            'идеиобразование',
                                            'общее развитие'
                                        ]"
                                    />
                                    <Select
                                        :rules="this.ruleUrgency"
                                        @inputed="getTaskUrgency"
                                        label="Срочность"
                                        :items="[
                                            'В первую очередь',
                                            'Обычная',
                                            'Похуй'
                                        ]"
                                    />
                                    <TextArea :rules="this.ruleDescription" label="Описание задачи" @customChange="getTaskDescription" />
                                    <TextField :rules="this.ruleAuthor" label="Автор задачи" @customChange="getTaskAuthor"/>

                                    <div class="mx-15">Примерная дата выполнения:</div>
                                    <v-responsive
                                        class="mx-auto"
                                        max-width="500"
                                    >
                                        <v-text-field :rules="this.ruleDateEnd" v-model="task.date_end" variant="solo-inverted"  type="date"></v-text-field>
                                    </v-responsive>
                                

                                
                            </v-card-text>

                            <v-card-actions>
                                <v-btn 
                                    color="green"
                                    variant="flat"
                                    text="добавить"
                                    type="submit"
                                >
                                </v-btn>
                                <v-spacer></v-spacer>

                                <v-btn
                                    variant="flat"
                                    color="red"
                                    text="закрыть"
                                    @click="isActive.value = false"
                                ></v-btn>
                            </v-card-actions>
                        </v-form>
                    </v-card>
                
                </template>
            </v-dialog>
            
        </v-container>
    </v-container>
</template>

<script setup>
import axios from 'axios';

const task = reactive({
    name: "",
    area: "",
    urgency: "",
    description: "",
    author: "",
    date_end: "",
});


const url_api = inject('url_api');
const router = useRoute();

function getTaskName(event){
    task.name = event;
}

function getTaskArea(event){
    task.area = event;
}

function getTaskUrgency(event){
    task.urgency = event;
}

function getTaskDescription(event){
    task.description = event;
}

function getTaskAuthor(event){
    task.author = event;
}

function getTaskDateEnd(event){
    task.date_end = event;
}

function submit(){
    // axios.post(`${url_api}tasks`, {
    //     name: task.name,
    //     area: task.area,
    //     urgency: task.urgency,
    //     description: task.description,
    //     author: task.author,
    //     date_end: task.date_end
    // });
    console.log('pizda')
}
</script>

<script>
  export default {

    data: () => ({
        dialog: false,
      ruleName: [
        value => {
          if (value) return true

          return 'Заполните "Название" задачи'
        },
      ],

      ruleArea: [
        value => {
            if (value) return true

            return 'Заполните "Область" задачи'
        }
      ],

      ruleUrgency: [
        value => {
            if (value) return true

            return 'Заполните "Срочность" задачи'
        }
      ],

      ruleDescription: [
        value => {
            if (value) return true

            return 'Заполните "Описание" задачи'
        }
      ],

      ruleAuthor: [
        value => {
            if (value) return true

            return 'Заполните "Автор" задачи'
        }
      ],

      ruleDateEnd: [
        value => {
            if (value) return true

            return 'Заполните "Дату выполнения" задачи'
        }
      ],
    }),
  }
</script>