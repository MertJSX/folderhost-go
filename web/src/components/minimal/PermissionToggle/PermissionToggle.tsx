import SwitchToggle from "../SwitchToggle/SwitchToggle"

interface PermissionToggleProps {
  label: string
  checked: boolean
  onChange: (checked: boolean) => void
}

const PermissionToggle = ({ label, checked, onChange }: PermissionToggleProps) => {
  return (
    <div className="flex items-center justify-between p-2 hover:bg-gray-500 rounded transition-colors">
      <span className="text-white text-sm">{label}</span>
      <SwitchToggle checked={checked} onChange={onChange} />
    </div>
  )
}

export default PermissionToggle;