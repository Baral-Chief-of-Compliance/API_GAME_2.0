<template>
  <div>
    <NuxtLayout>
      <Header  />
      <NuxtPage/>
    </NuxtLayout>
  </div>
</template>

<script>


export default{

  setup(){
    provide('basicColor', '#b7410e');
    provide('additionalColor', '#696969');
    provide('url_api', 'http://127.0.0.1:5000/the_stone_sword/api/v1.0/')
  }
}
</script>

<style>

*{
  font-family: "better-vcr";
}

@font-face {
  font-family: "better-vcr";
  src: url("./assets/fonts/better-vcr_0.ttf");
}
</style>