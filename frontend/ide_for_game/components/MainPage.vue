<template>
    <v-container>
        <TitlePage :title="title" />
        <v-container class="mx-15 d-flex mt-5">
            <v-row>
                <v-col cols="12" md="3">
                    <v-btn class="mx-5" width="300" height="200" variant="outlined"   
                        @click="this.$router.push({ path: path+'/create'})"
                        :color="additionalColor">
                        {{ title_for_batton }}
                    </v-btn>
                </v-col>

                <v-col cols="12" md="3" v-for="d in data">
                    <v-hover>
                        <template v-slot:default="{ isHovering, props }">
                            <v-card  
                                class="mx-5 d-flex align-center justify-center"  
                                width="300" 
                                height="200"
                                v-bind="props"
                                :color="isHovering ? additionalColor : undefined"
                                @click="this.$router.push({ path: path+`/${d.id}`})"
                            >
                                <v-card-title>
                                    {{ d.name }} 
                                </v-card-title>
                            </v-card>
                        </template>
                    </v-hover>
                </v-col>
            </v-row>
        </v-container>
    </v-container>
</template>

<script>

    export default{
    props: {
        title: String,
        data: Array,
        basicColor: String,
        additionalColor: String,
        title_for_batton: String,
        type_menu: String,
        path: String
    },
    setup() {
        const basicColor = inject('basicColor');
        const additionalColor = inject('additionalColor');
        return { basicColor, additionalColor };
    }
}
</script>