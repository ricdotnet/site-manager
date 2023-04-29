<template>
  <input type="file" @change="onInputChange">
  <div ref="outer" class="relative w-[500px]">
<!--    <img class="w-[500px] h-auto" ref="image" alt="original image"/>-->
    <canvas class="absolute top-0" ref="original" width="500" height="500"></canvas>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, reactive, ref } from "vue";

  const image = ref<HTMLImageElement>();
  const original = ref<HTMLCanvasElement>();
  const outer = ref<HTMLDivElement>();

  const state = reactive<{
    isDrawing: boolean,
    startX: number,
    startY: number,
    ctx: null | CanvasRenderingContext2D,
  }>({
    isDrawing: false,
    startX: 0,
    startY: 0,
    ctx: null,
  });

  onMounted(() => {
    state.ctx = original.value!.getContext('2d');

    original.value!.addEventListener('mousedown', e => {
      state.isDrawing = true;
      state.startX = e.offsetX;
      state.startY = e.offsetY;
    });

    original.value!.addEventListener('mousemove', e => {
      if (state.isDrawing) {
        const currentX = e.offsetX;
        const currentY = e.offsetY;

        // Calculate the width of the box
        // const boxWidth = Math.min(currentX - state.startX, original.value!.width - state.startX);
        //
        // // Clear the canvas
        // state.ctx!.clearRect(0, 0, original.value!.width, original.value!.height);
        //
        // // Draw the box
        // state.ctx!.strokeRect(state.startX, state.startY, boxWidth, currentY - state.startY);
        // Calculate the size of the square
        const size = Math.min(Math.abs(currentX - state.startX), Math.abs(currentY - state.startY));

        // Clear the canvas
        state.ctx!.clearRect(0, 0, original.value!.width, original.value!.height);

        // Draw the square
        if (currentX < state.startX && currentY < state.startY) {
          // Top-left quadrant
          state.ctx!.strokeRect(state.startX - size, state.startY - size, size, size);
        } else if (currentX < state.startX && currentY > state.startY) {
          // Bottom-left quadrant
          state.ctx!.strokeRect(state.startX - size, state.startY, size, size);
        } else if (currentX > state.startX && currentY < state.startY) {
          // Top-right quadrant
          state.ctx!.strokeRect(state.startX, state.startY - size, size, size);
        } else {
          // Bottom-right quadrant
          state.ctx!.strokeRect(state.startX, state.startY, size, size);
        }
      }
    });

    original.value!.addEventListener('mouseup', e => {
      state.isDrawing = false;
    });
  });

  function onInputChange(e: Event) {
    if ('target' in e && 'files' in e.target) {
      const img = new Image();
      img.onload = () => {

        // if (image.value && 'src' in image.value) {
        //   image.value.src = img.src;
        // }
        //
        // if (original && 'value' in original) {
        //   original.value.width = outer.value?.offsetWidth;
        //   original.value.height = outer.value?.offsetHeight;
        //   debugger;
        // }

        const canvasWidth = original.value!.width;
        const canvasHeight = original.value!.height;
        const imageWidth = img.width;
        const imageHeight = img.height;
        const aspectRatio = imageWidth / imageHeight;
        let drawWidth = canvasWidth;
        let drawHeight = drawWidth / aspectRatio;
        if (drawHeight > canvasHeight) {
          drawHeight = canvasHeight;
          drawWidth = drawHeight * aspectRatio;
        }
        const x = (canvasWidth - drawWidth) / 2;
        const y = (canvasHeight - drawHeight) / 2;
        state.ctx!.drawImage(img, x, y, drawWidth, drawHeight);

      };
      img.src = URL.createObjectURL(e.target.files[0]);
    }
  }
</script>
