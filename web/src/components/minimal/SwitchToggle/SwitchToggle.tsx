import React from 'react';
import "./switchtoggle.css";

interface SwitchToggleProps {
    checked: boolean,
    onChange: (checked: boolean) => void
}

const SwitchToggle: React.FC<SwitchToggleProps> = ({checked, onChange}) => {
  return (
    <label className='switch-toggle'>
        <input type="checkbox" checked={checked} onChange={() => {
            onChange(!checked)
        }} />
        <span className='slider' />
    </label>
  )
}

export default SwitchToggle