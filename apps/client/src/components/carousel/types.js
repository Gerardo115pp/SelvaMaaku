/**
 * @typedef {Object} CarouselItemParams
 * @property {number} id
 * @property {number} width
 * @property {number} height
 * @property {string} src
*/

export class CarouselItem {
    /**
     * @param {CarouselItemParams} params
     */
    constructor(params) {
        this.id = params.id;
        this.src = params.src;
        this.width = params.width || 0;
        this.height = params.height || 0;
    }
}