<script>
    import { browser } from "$app/environment";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import About from "@pages/Home/sections/About.svelte";
    import Amenities from "@pages/Home/sections/Amenities.svelte";
    import Banner from "@pages/Home/sections/Banner.svelte";
    import Hero from "@pages/Home/sections/Hero.svelte";
    import HouseModels from "@pages/Home/sections/HouseModels.svelte";
    import Location from "@pages/Home/sections/Location.svelte";
    import QuoteBanner from "@pages/Home/sections/QuoteBanner.svelte";
    import { onMount, onDestroy } from "svelte";
    import { cleanViewportObserver } from "@components/viewport_actions/useViewportActions";


    onMount(() => {
        setGradientType();
    });

    onDestroy(() => {
        if (browser) {
            cleanViewportObserver();
        }
    })
    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        const setGradientType = () => {
            if (!browser) return;
            
            const root_element = document.querySelector('html');

            if (root_element) {
                root_element.setAttribute('data-gradient-type', 'main-page');
            }
        }
    
    /*=====  End of Methods  ======*/
    
    
</script>

<main id="smk-home-page">
    <Hero />
    <About />
    <Location />
    <Banner />
    <HouseModels />
    <Amenities />
    <QuoteBanner /> 
</main>