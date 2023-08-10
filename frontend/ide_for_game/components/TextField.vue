<template>
    <v-responsive
        class="mx-auto"
        max-width="500"
    >
        <v-text-field
            :label="label"
            variant="outlined"
            @change="$emit('customChange', $event.target.value)"
        >
        </v-text-field>

    </v-responsive>
</template>

<script>
    export default{
        props:{
            label: String
        }
    }
</script>