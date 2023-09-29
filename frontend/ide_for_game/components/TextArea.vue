<template>

    <v-responsive
        class="mx-auto"
        max-width="500"
    >
        <v-textarea 
            :label="label"
            :rules="rules" 
            rows="8"
            variant="solo-inverted"
            @change="$emit('customChange', $event.target.value)"
        >
            
        </v-textarea>

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