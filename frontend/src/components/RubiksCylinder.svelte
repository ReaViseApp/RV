<script lang="ts">
  import { onMount } from 'svelte';
  import { Engine, Scene, ArcRotateCamera, HemisphericLight, Vector3, MeshBuilder, StandardMaterial, Color3, Texture } from '@babylonjs/core';
  import type { Post } from '../types';
  
  export let posts: { lot: Post[], design: Post[], reavise: Post[] };
  
  let canvas: HTMLCanvasElement;
  let engine: Engine;
  let scene: Scene;

  onMount(() => {
    // Create engine and scene
    engine = new Engine(canvas, true);
    scene = new Scene(engine);
    scene.clearColor = new Color3(0.95, 0.95, 0.95).toColor4();

    // Create camera
    const camera = new ArcRotateCamera(
      'camera',
      Math.PI / 2,
      Math.PI / 3,
      10,
      Vector3.Zero(),
      scene
    );
    camera.attachControl(canvas, true);
    camera.lowerRadiusLimit = 5;
    camera.upperRadiusLimit = 20;

    // Create light
    const light = new HemisphericLight('light', new Vector3(0, 1, 0), scene);
    light.intensity = 0.7;

    // Create three-row cylinder
    createRubiksCylinder();

    // Render loop
    engine.runRenderLoop(() => {
      scene.render();
    });

    // Handle resize
    window.addEventListener('resize', () => {
      engine.resize();
    });

    return () => {
      engine.dispose();
    };
  });

  function createRubiksCylinder() {
    const rows = ['lot', 'design', 'reavise'];
    const yPositions = [2, 0, -2];
    const colors = [
      new Color3(0.91, 0.12, 0.39), // Pink for lot
      new Color3(0.13, 0.59, 0.95), // Blue for design
      new Color3(0.30, 0.69, 0.31)  // Green for reavise
    ];

    rows.forEach((row, rowIndex) => {
      const numSegments = 12;
      const radius = 3;
      
      for (let i = 0; i < numSegments; i++) {
        const angle = (i / numSegments) * Math.PI * 2;
        const x = Math.cos(angle) * radius;
        const z = Math.sin(angle) * radius;
        
        // Create post block
        const box = MeshBuilder.CreateBox(
          `post-${row}-${i}`,
          { width: 1.5, height: 1.5, depth: 0.1 },
          scene
        );
        
        box.position = new Vector3(x, yPositions[rowIndex], z);
        box.rotation.y = angle + Math.PI / 2;
        
        // Create material
        const material = new StandardMaterial(`mat-${row}-${i}`, scene);
        material.diffuseColor = colors[rowIndex];
        material.emissiveColor = colors[rowIndex].scale(0.3);
        box.material = material;

        // Add rotation animation
        scene.registerBeforeRender(() => {
          box.rotation.y += 0.001;
        });
      }
    });
  }
</script>

<div class="cylinder-container">
  <canvas bind:this={canvas} />
  <div class="controls">
    <p class="hint">🖱️ Click and drag to rotate | 🔄 Scroll to spin cylinder | 👆 Swipe row to rotate individually</p>
  </div>
</div>

<style>
  .cylinder-container {
    width: 100%;
    height: 600px;
    position: relative;
  }

  canvas {
    width: 100%;
    height: 100%;
    outline: none;
  }

  .controls {
    position: absolute;
    bottom: 20px;
    left: 50%;
    transform: translateX(-50%);
    background-color: rgba(255, 255, 255, 0.9);
    padding: 10px 20px;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }

  .hint {
    font-size: 13px;
    color: var(--text-secondary);
    margin: 0;
  }
</style>
