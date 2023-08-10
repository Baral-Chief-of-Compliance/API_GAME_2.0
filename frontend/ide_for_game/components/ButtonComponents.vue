<template>

    <v-responsive
        class="mx-auto"
        max-width="500"
    >
        <v-btn 
        width="467"
        size="x-large"
            v-on="$listeners" 
            :color="basicColor"
        >
            <span class="text-white">{{ label }}</span>
        </v-btn>
    </v-responsive>

</template>

<script>
export default{
    setup(){
        const basicColor = inject('basicColor')

        return { basicColor }
    },

    props: {
        label: String
    }
}
</script>