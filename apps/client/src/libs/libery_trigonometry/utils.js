
export const degToRad = deg => {
    return (deg * Math.PI) / 180;
}

export const getTrianglePoints = (cx, cy, d, angle_offset=-90) => {
    const radius = d / Math.sqrt(3);

    /**
     * @typedef {Object} Point
     * @property {number} x
     * @property {number} y
     */
    
     let v1 = {
        x: cx + radius * Math.cos(degToRad(angle_offset)),
        y: cy + radius * Math.sin(degToRad(angle_offset))
     }

     let v2 = {
        x: cx + radius * Math.cos(degToRad(angle_offset + 120)),
        y: cy + radius * Math.sin(degToRad(angle_offset + 120))
     }

     let v3 = {
        x: cx + radius * Math.cos(degToRad(angle_offset + 240)),
        y: cy + radius * Math.sin(degToRad(angle_offset + 240))
     }

     return `M${v1.x} ${v1.y}L${v2.x} ${v2.y}L${v3.x} ${v3.y}Z` 
}