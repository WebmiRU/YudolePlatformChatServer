<script>
import axios from "axios";
import {config} from "./store/config.js";

export default {
  data() {
    return {

    }
  },
  methods: {
    configGet() {
      axios.get('http://127.0.0.1:5800/config', {}).then((r) => {
        Object.assign(config, r.data)
      }).catch((e) => {
        // @TODO Error
        console.log(e)
      })
    },
    configPost() {
      axios.post('http://127.0.0.1:5800/config', config)
          .then((r) => {
            console.log(r)
          })
          .catch((e) => {
            // @TODO Error
            console.log(e)
          })
    },
  },
  beforeMount() {
    this.configGet()
    console.log('STORE', config)
  }
}
</script>

<template>
  <header>
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" href="https://bulma.io">
          <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">
        </a>

        <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item">
            Home
          </a>

          <a class="navbar-item">
            Documentation
          </a>

          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link">
              More
            </a>

            <div class="navbar-dropdown">
              <a class="navbar-item">
                About
              </a>
              <a class="navbar-item">
                Jobs
              </a>
              <a class="navbar-item">
                Contact
              </a>
              <hr class="navbar-divider">
              <a class="navbar-item">
                Report an issue
              </a>
            </div>
          </div>
        </div>

        <div class="navbar-end">
<!--          <div class="navbar-item">-->
<!--            <label class="checkbox">-->
<!--              <input type="checkbox"> Режим эксперта</label>-->
<!--          </div>-->
          <div class="navbar-item">
            <div class="buttons">
              <button @click="configPost" class="button is-success">Сохранить</button>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </header>

  <div class="columns">
  <aside class="menu column is-4" style="padding: 30px">
    <p class="menu-label">
      Чаты
    </p>
    <ul class="menu-list">
<!--      <a class="is-active">GoodGame</a>-->
      <li><router-link :to="{name: 'chats.goodgame'}" active-class="is-active">GoodGame</router-link></li>
      <li><router-link :to="{name: 'chats.trovo'}" active-class="is-active">Trovo</router-link></li>
      <li><router-link :to="{name: 'chats.twitch'}" active-class="is-active">Twitch</router-link></li>
      <li><router-link :to="{name: 'chats.youtube'}" active-class="is-active">YouTube</router-link></li>
    </ul>
    <p class="menu-label">
      Прочие параметры
    </p>
    <ul class="menu-list">
      <li><a>Team Settings</a></li>
      <li><a>Manage Your Team</a></li>
    </ul>

  </aside>
  <main class="column is-8" style="padding: 30px">
    <router-view></router-view>
  </main>
  </div>
</template>
