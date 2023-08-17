<script>
  import logo from './assets/images/image.png'
  import {Lint} from '../wailsjs/go/main/App.js'

  let greetText = "Please upload your file (.txt or .md), at least 200 words";
  let specText = "Indicate data as singular (1) or plural (2).";
  let spec;
  let resultText = []; 

  function lint() { 
    Lint(spec).then(result => resultText = result)
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div class="greeting" id="greeting">{greetText}</div>
  <div class="specification" id="specification">{specText}</div>

  <!-- Input from user -->
  <input autocomplete="off" bind:value={spec} class="input" id="name" type="text"/>
  <button class="btn" on:click={lint}>Lint it</button>
  
  <!-- Lists out each line from output -->
  <ul>
    {#each resultText as res}
    <li>{res}</li>
    {/each}
     </ul>
</main>

<style>
  #logo {
    display: block;
    width: 20%;
    height: 20%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }
</style>
