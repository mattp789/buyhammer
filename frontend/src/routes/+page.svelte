<style>
  body {
    background-color: #333;
    color: white;
    display: grid;
    place-items: center; 
    margin: 0;
    }
  textarea {
    height: 200px;
    resize: vertical;
    background-color: #666;
    color: white;
    border-radius: 5px;
    width: 300px;
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
    margin-top: 10px;
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
    }

    @keyframes spin {
        0% { transform: rotate(0deg); }
        100% { transform: rotate(360deg); }
    }
</style>



<script lang="ts">
import axios from 'axios';
let multiplier = 100;
let responseData = '';
let data = "";
let loading = false;
async function listData(){
    responseData = '';
    loading = true;
    // const dataitems = data.split(('\n'))
    const lines = data.split('\n');
    const dataitems = lines.filter(line => 
    /\(\d+\)/.test(line) && !line.startsWith('App:')
    ).map(line => line.replace(/\s\(\d+\)/, ''));

    try {
          const response = await axios({
          method: 'post',
          url: 'http://localhost:8080/calculate',
          data: {
              dataitems,
              multiplier
          }
          });
      responseData = response.data.message;
      loading = false;
    } catch(error){
      loading = false;

      console.error('Error while fetching data:', error);

      if (axios.isAxiosError(error)) {
          responseData = error.response?.data?.message || 'An error occurred while processing the request.';
      } else {
          responseData = 'An unexpected error occurred.';
      }
  }
}
</script>

<body>
<h1>Welcome to Buyhammer!</h1>
<h3>Copy paste your list text in the box below to calculate the MSRP value.</h3>
<h4>Note: This app cannot currently calculate reinforced units, so please enter each unit individually.</h4>
<div>
  <textarea placeholder="Paste your list here..." bind:value={data}></textarea>
</div>
<label>
  % of MSRP
	<input type="number" bind:value={multiplier} min="0" max="100" />
	<input type="range" bind:value={multiplier} min="0" max="100" />
</label>
{#if loading}
    <div class="loading-spinner"></div>
{/if}
{#if responseData}
    <div>
        <p>{responseData}</p>
    </div>
{/if}
<div>
  <button on:click={listData}>Submit</button>
</div>
</body>
