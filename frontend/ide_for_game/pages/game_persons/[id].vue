<template>
    <v-container>
        <TitlePage title="Игровой персонаж: " />
        <Tubs 
            :tabs="['Информация', 'Характеристики','Иконка', 'Анимации']" 
            @get-index="updateIndex"
        />
    </v-container>

   
    
</template>

<script setup>

    const tab = ref("sdsdsdsd")

    function updateIndex(index){
        tab.value = index
    }
</script>