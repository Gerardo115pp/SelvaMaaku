<script>
    import viewport from "@components/viewport_actions/useViewportActions";
    
    export let id = `lazy-wrapper-${crypto.randomUUID()}`;
    export let className;
    export let image_url;

    let image_element;
    let image_src = "";

    /** @type {boolean} whether or not the image src has been loaded */
    let is_loaded = false;

    /** @type {boolean} whether or not the image src has entered the viewport */
    let entered_viewport = false;

    $: if(image_src !== image_url && entered_viewport) {
        is_loaded = false;
        mountSrc();
    }

    const handleViewportEnter = e => {
        if (entered_viewport && is_loaded) return;

        entered_viewport = true;
        mountSrc();
    };

    const mountSrc = () => {
        image_element = document.createElement('img');

        image_element.onload = () => {
            image_src = image_element.src;
            is_loaded = true;
        };

        image_element.src = image_url;
    }
</script>

<div use:viewport on:viewportEnter={handleViewportEnter}  class={`lazy-wrapper ${className !== undefined ? className : ''}`} {id}>
    {#if is_loaded}
        <slot name="lazy-wrapper-image" {image_src}></slot>
    {/if}
</div>

<style>
    
    @keyframes placeholderShimmer {
        0% {
            background-position: 220% 0
        }
        100% {
            background-position:  -210% 0
        }
    }
    
    .lazy-wrapper:empty {
        animation-duration: 1s;
        animation-fill-mode: forwards;
        animation-iteration-count: infinite;
        animation-name: placeholderShimmer;
        animation-timing-function: linear;
        width: 100%;
        height: 100%;
        background: rgb(236,236,236);
        background: linear-gradient(98deg, var(--clear-2) 34%, var(--clear-4) 38%, var(--clear-4) 50%, var(--clear-2) 71%); 
        background-size: 130%;
    }
</style>