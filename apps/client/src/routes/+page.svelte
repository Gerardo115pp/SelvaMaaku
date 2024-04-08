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
    import { onMount } from "svelte";


    onMount(() => {
        setGradientType();
    });
    
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
    <div class="wixarica-underlay" class:adebug={false}>
        <WixaricaIcon 
            opacity={0.1}
        />
    </div>
    <Hero />
    <About />
    <Location />
    <Banner />
    <HouseModels />
    <Amenities />
    <QuoteBanner /> 
</main>

<style>
    main#smk-home-page {
        position: relative;

        & .wixarica-underlay {
            position: absolute;
            top: -3%;
            left: 50%;
            width: 84%;
            max-width: var(--design-content-width);
            z-index: var(--z-index-b-3);
            transform: translateX(-50%);
        }
    }
    
</style>