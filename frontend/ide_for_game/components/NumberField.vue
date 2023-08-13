<template>
    <v-responsive
        class="mx-auto"
        max-width="500"
    >
        <v-text-field
            :label="label"
            variant="outlined"
            v-model="field_number"
            @change="$emit('customChange', $event.target.value)"
            type="number"
            :rules="[
                field_number > -1 && field_number < 11 || 'Доверие от 0 до 10'
            ]"
        >
        </v-text-field>

    </v-responsive>

</template>

<script>
    export default{
        data(){
            return{
                field_number: 0
            }
        },
        props:{
            label: String
        }
    }
</script>