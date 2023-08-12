<template>
    <v-container class="mx-10">
        <v-row>
            <v-col cols="2" v-for="(item, index) in tabs">
                <v-hover>
                    <template v-slot:default="{ isHovering, props }">
                        <v-card 
                            class="d-flex align-center justify-center"
                            v-bind="props"
                            :color="(tab == index) ? 'indigo'  : (isHovering) ? additionalColor : 'undefind' "
                            @click="selectTab(index)"
                        >
                            <v-card-title>
                                {{ item }}
                            </v-card-title>
                        </v-card>
                    </template>
                </v-hover>

            </v-col>
        </v-row>
    </v-container>

</template>

<script setup>


const props = defineProps({
    tabs: Array,
    tab: Number
})

const emit = defineEmits(["getIndex"])


function selectTab(index){
    emit("getIndex", index)
}

const additionalColor = inject("additionalColor")
</script>                            
<!-- :color=" isHovering ? additionalColor : undefined" -->