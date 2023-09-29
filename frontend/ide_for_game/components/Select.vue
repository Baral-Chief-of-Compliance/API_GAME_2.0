<template>
    <v-responsive
        class="mx-auto"
        max-width="500"
    >
        <v-select
            :rules="rules"
            :label="label"
            :items="items"
            v-model="selected"
            variant="solo-inverted"
        >   

        </v-select>
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
            items: Array,
            rules: {
                type: Array,
                required: false
            }
        }
    }
</script>

