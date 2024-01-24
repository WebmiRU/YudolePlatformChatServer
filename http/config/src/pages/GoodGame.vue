<script setup>
import {store} from "../store.js";
import Trovo from "./Trovo.vue";
</script>

<template>
  <h1 class="title is-1">GoodGame</h1>
  <div class="field">
    <label class="label">Имя пользователя</label>
    <div class="control">
      <input class="input" type="text" placeholder="Имя пользователя">
    </div>
  </div>

  <div class="field">
    <label class="label">Пароль</label>
    <div class="control">
      <input class="input" type="text" placeholder="Пароль">
    </div>
  </div>


  <div class="field">
    <label class="label">Каналы</label>
  </div>

  <div v-for="(v1, i) in store.services.goodgame" class="field has-addons">
    <div class="control is-expanded">
      <input v-model="store.services.goodgame[i]" class="input" type="number" min="1" max="999999999" placeholder="Номер канала">
    </div>
    <div class="control" v-if="store.services.goodgame.length > 1">
      <a @click="store.services.goodgame.splice(i, 1)" class="button is-danger">-</a>
    </div>
    <div class="control">
      <a @click="store.services.goodgame.push(null)" class="button is-primary">+</a>
    </div>
  </div>
</template>
