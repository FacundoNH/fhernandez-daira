<script>
    import { onMount } from 'svelte';
    
    let num1 = '';
    let num2 = '';
    let operador = '';
    let resultado = '';
    let historial = [];
  
    // Función para realizar la operación en el backend
    const operar = async () => {
      const response = await fetch('http://localhost:8080/api/operar', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          num1: parseInt(num1),
          num2: parseInt(num2),
          operador,
        }),
      });
  
      const data = await response.json();
      resultado = data.resultado;
  
      // Actualizar historial después de cada operación
      actualizarHistorial();
    };
  
    // Función para obtener el historial desde el backend
    const obtenerHistorial = async () => {
      const response = await fetch('http://localhost:8080/api/historial');
      historial = await response.json();
    };
  
    // Función para actualizar el historial
    const actualizarHistorial = async () => {
      await obtenerHistorial();
    };
  
    // Llamar a obtenerHistorial al cargar la página
    onMount(obtenerHistorial);
  </script>
  
  <main>
    <section>
      <h2>Calculadora</h2>
      <input bind:value={num1} type="number" placeholder="Número 1" />
      <input bind:value={num2} type="number" placeholder="Número 2" />
  
      <div>
        <button on:click={() => { operador = '+'; operar(); }}>+</button>
        <button on:click={() => { operador = '-'; operar(); }}>-</button>
        <button on:click={() => { operador = '*'; operar(); }}>*</button>
        <button on:click={() => { operador = '/'; operar(); }}>/</button>
      </div>
  
      <button on:click={operar}>Calcular</button>
    </section>
  
    <section>
      <h2>Resultado</h2>
      <p>{resultado}</p>
    </section>
  
    <section>
      <h2>Historial</h2>
      {#each historial as operacion (operacion.Num1)}
        <p>{operacion.Num1} {operacion.Operador} {operacion.Num2} = {operacion.Resultado}</p>
      {/each}
    </section>
  </main>
  
  <style>
    /* styles.css */

main {
  display: flex;
  justify-content: space-around;
  align-items: center;
  height: 100vh;
}

section {
  background-color: #f0f0f0;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 300px;
  text-align: center;
}

input {
  margin: 10px 0;
  padding: 8px;
  width: 80%;
  border: 1px solid #ccc;
  border-radius: 5px;
  outline: none;
}

button {
  margin: 5px;
  padding: 10px 20px;
  background-color: #4caf50;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #45a049;
}

h2 {
  color: #333;
}

p {
  margin: 5px 0;
  color: #777;
}
  </style>
  