<template>

    <v-responsive
        class="mx-auto"
        max-width="500"
    >
        <v-textarea 
            :label="label" 
            rows="20"
            variant="outlined"
            @change="$emit('customChange', $event.target.value)"
        >
            
        </v-textarea>

    </v-responsive>
</template>

<script>
    export default{
        props:{
            label: String
        }
    }
</script>