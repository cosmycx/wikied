import React from 'react';

export default ({viewWidths, setViewWidths}) => {
  const handleDrag = function(event) {
    const available = 98;
    const screenWidth = window.outerWidth;
    const mousePosition = event.clientX;
    if(mousePosition === 0) return;
    const left = Math.round(mousePosition/screenWidth * 100);
    setViewWidths({left, right: available - left});
  }
  return (
  <div onDrag={handleDrag} className="editor__vertical-handler"></div>
)};