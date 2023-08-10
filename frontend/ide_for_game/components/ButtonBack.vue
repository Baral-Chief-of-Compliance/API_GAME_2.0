<template>
    <v-btn
        color="black"
        variant="outlined"
        width="300"
        v-on="$listeners" 
        class="mx-15 my-5"
    >
        <Icon name="cil:arrow-thick-to-left" class="mr-3"></Icon> {{ title }}
    </v-btn>
</template>

<script>
export default{
    props: {
        title: String
    }
}
</script>