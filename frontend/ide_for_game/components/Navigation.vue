<template>
      <v-layout>
        <v-navigation-drawer
          expand-on-hover
          rail
          :color="basicColor"
          class="text-white"
        >
          <v-list>
            <v-list-item
              title="The Stone Sword"  
            >
                <template v-slot:prepend>
                    <Icon name="fa6-solid:hammer" class="mr-3" color="white"></Icon>
                </template>

                <template v-slot:subtitle>
                    <div class="subtitle">Рабочее пространство</div>
                </template>
            </v-list-item>
          </v-list>
  
          <v-divider></v-divider>
  
          <v-list density="compact" nav>
            <NuxtLink to="/">
                <v-list-item title="Главная" value="index">
                    <template v-slot:prepend>
                        <NavigationIcon icon_name="fa6-solid:book"/>
                    </template>
                </v-list-item>
            </NuxtLink>

            <NuxtLink to="/stories">
                <v-list-item title="Сюжет" value="story" >
                    <template v-slot:prepend>
                        <NavigationIcon icon_name="fa6-solid:book-open"/>
                    </template>
                </v-list-item>
            </NuxtLink>
            

            <NuxtLink to="/game_persons">
                <v-list-item title="Персонажи" value="game_persons">
                    <template v-slot:prepend>
                        <NavigationIcon icon_name="fa6-solid:person"/>
                    </template>
                </v-list-item>
            </NuxtLink>

            <NuxtLink to="/factions">
                <v-list-item prepend-icon="mdi-star" title="Фракции" value="factions">
                    <template v-slot:prepend>
                        <NavigationIcon icon_name="fa6-solid:flag"/>
                    </template>
                </v-list-item>
            </NuxtLink>
            
            
          </v-list>
        </v-navigation-drawer>

      </v-layout>
  </template>

<script>
    export default{
        setup(){
            const basicColor = inject('basicColor')

            return { basicColor }
        }
    }
</script>

<style scoped>

    .subtitle{
        color: white;
        font-size: 10px;
    }

    a {
        color: white;
        text-decoration: none;
    }

</style>