<template>
    <v-dialog>
        <template v-slot:activator="{ props }">
            <v-btn 
                    v-model="dialog"
                    color="indigo"
                    block
                    class="mt-15"
                    v-bind="props"
                >
                    {{ title }}
            </v-btn> 
        </template>

        <v-card>
            <v-card-text>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
            </v-card-text>
            <v-card-actions>
            <v-btn color="primary" block @click="dialog.status = false">Close Dialog</v-btn>
            </v-card-actions>
        </v-card>

    </v-dialog>
</template>

<script setup>

    const dialog = reactive({
        status: false
    })

    const props = defineProps(
        {
            title: String
        }
    )

</script>