<template>
    <v-responsive
        class="mx-auto"
        max-width="500"
    >
        <v-text-field
            :rules="rules"
            :label="label"
            variant="solo-inverted"
            @change="$emit('customChange', $event.target.value)"
        >
        </v-text-field>

    </v-responsive>
</template>

<script>
    export default{
        props:{
            label: String,
            rules: {
                type: Array,
                required: false
            }
        }
    }
</script>