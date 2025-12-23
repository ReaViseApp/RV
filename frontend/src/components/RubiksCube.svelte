<script lang="ts">
  import { onMount } from 'svelte';
  import { Engine, Scene, ArcRotateCamera, HemisphericLight, Vector3, MeshBuilder, StandardMaterial, Color3 } from '@babylonjs/core';
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
      Math.PI / 4,
      Math.PI / 4,
      12,
      Vector3.Zero(),
      scene
    );
    camera.attachControl(canvas, true);
    camera.lowerRadiusLimit = 8;
    camera.upperRadiusLimit = 20;

    // Create light
    const light = new HemisphericLight('light', new Vector3(0, 1, 0), scene);
    light.intensity = 0.7;

    // Create Rubik's Cube
    createRubiksCube();

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

  function createRubiksCube() {
    const categories = ['lot', 'design', 'reavise'];
    const colors = [
      new Color3(0.91, 0.12, 0.39), // Pink for lot
      new Color3(0.13, 0.59, 0.95), // Blue for design
      new Color3(0.30, 0.69, 0.31)  // Green for reavise
    ];

    // Create 3x3x3 cube structure
    for (let x = -1; x <= 1; x++) {
      for (let y = -1; y <= 1; y++) {
        for (let z = -1; z <= 1; z++) {
          // Skip center cube
          if (x === 0 && y === 0 && z === 0) continue;

          const box = MeshBuilder.CreateBox(
            `cube-${x}-${y}-${z}`,
            { size: 0.95 },
            scene
          );
          
          box.position = new Vector3(x * 1.05, y * 1.05, z * 1.05);
          
          // Assign color based on x position (column)
          const categoryIndex = x + 1; // -1, 0, 1 -> 0, 1, 2
          const material = new StandardMaterial(`mat-${x}-${y}-${z}`, scene);
          material.diffuseColor = colors[categoryIndex];
          material.emissiveColor = colors[categoryIndex].scale(0.2);
          box.material = material;
        }
      }
    }

    // Add edge labels
    createEdgeLabels();
  }

  function createEdgeLabels() {
    const labels = [
      { text: 'The Lot', position: new Vector3(-1.5, 2, 0), color: new Color3(0.91, 0.12, 0.39) },
      { text: 'Design', position: new Vector3(0, 2, 0), color: new Color3(0.13, 0.59, 0.95) },
      { text: 'ReaVise', position: new Vector3(1.5, 2, 0), color: new Color3(0.30, 0.69, 0.31) }
    ];

    labels.forEach(label => {
      const plane = MeshBuilder.CreatePlane(
        `label-${label.text}`,
        { width: 1, height: 0.3 },
        scene
      );
      plane.position = label.position;
      
      const material = new StandardMaterial(`label-mat-${label.text}`, scene);
      material.diffuseColor = label.color;
      material.emissiveColor = label.color.scale(0.5);
      plane.material = material;
    });
  }
</script>

<div class="cube-container">
  <canvas bind:this={canvas} />
  <div class="controls">
    <p class="hint">🖱️ Click and drag to rotate | 📏 Cmd+↑/↓ to resize | 👆 Swipe column to rotate</p>
  </div>
</div>

<style>
  .cube-container {
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
