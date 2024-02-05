<script>
	import Calculadora from '/components/Calculadora.svelte';
  </script>
  
  <main>
	<h1>Calculadora</h1>
	<Calculadora />
  </main>
  
  <style>
	main {
	  text-align: center;
	  padding: 1em;
	  max-width: 240px;
	  margin: 0 auto;
	}
  
	h1 {
	  color: #ff3e00;
	  text-transform: uppercase;
	  font-size: 2em; /* Ajustado el tamaño del título */
	  font-weight: 700; /* Haciendo el título más audaz */
	  margin-bottom: 1em; /* Agregado margen inferior para separación */
	}
  
	@media (min-width: 640px) {
	  main {
		max-width: none;
	  }
	}
  </style>
  