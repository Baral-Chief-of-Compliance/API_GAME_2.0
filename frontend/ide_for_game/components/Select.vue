<template>
    <v-responsive
        class="mx-auto"
        max-width="500"
    >
        <v-select
            :label="label"
            :items="items"
            v-model="selected"
            variant="outlined"
        >   

        </v-select>
        {{ selected }}
    </v-responsive>
</template>

<script setup>
    let selected = ref("");

    const emit = defineEmits(['inputed'])
        emit("inputed", selected);

</script>


<script>
    export default{
        props: {
            label: String,
            items: Array
        }
    }
</script>

