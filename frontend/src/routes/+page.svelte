<style>
  body {
    background-color: #333;
    color: white;
    }
  textarea {
    width: 30%;
    height: 200px;
    resize: vertical;
    background-color: #666;
    color: white;
    border-radius: 5px;
  }
  textarea::placeholder {
    color: white;
  }
  button {
    padding: 10px 20px;
    background-color: #007BFF;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }

  button:hover {
    background-color: #0056b3;
  }
  p {
    font-family: 'Montserrat', sans-serif; 
    font-size: 2rem; 
    color: white; 
    text-shadow: 2px 2px 8px rgba(0, 0, 0, 0.5); 
    text-align: match-parent; 
  }
  .loading-spinner {
        border: 8px solid #f3f3f3; 
        border-top: 8px solid #3498db; 
        border-radius: 50%;
        width: 40px;
        height: 40px;
        animation: spin 1s linear infinite;
        margin-top: 10px;
        margin-left: 12%;
    }

    @keyframes spin {
        0% { transform: rotate(0deg); }
        100% { transform: rotate(360deg); }
    }
</style>



<script lang="ts">
import axios from 'axios';
import {onMount} from 'svelte';
let responseData = '';
let data = "";
let loading = false;
async function listData(){
    responseData = '';
    loading = true;
    const dataitems = data.split('\n');
    const response = await axios({
        method: 'post',
        url: 'http://localhost:8080/calculate',
        data: {
            dataitems
        }
        });
    responseData = response.data.message;
    loading = false;
}
</script>

<body>
<div>
  <textarea placeholder="Type your units here..." bind:value={data}></textarea>
</div>
<button on:click={listData}>Submit</button>
{#if loading}
    <div class="loading-spinner"></div>
{/if}
{#if responseData}
    <div>
        <p>{responseData}</p>
    </div>
{/if}
</body>
