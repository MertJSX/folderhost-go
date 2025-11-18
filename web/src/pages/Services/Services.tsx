import { MdMiscellaneousServices } from "react-icons/md";
import { FaClock, FaTools } from "react-icons/fa";
import { BiSolidParty } from "react-icons/bi";
import { useEffect } from "react";

const Services = () => {
  useEffect(() => {
    document.title = "Services - folderhost"
  },[])
  return (
    <div className="flex items-center justify-center min-h-[calc(100vh-120px)] bg-slate-900 p-6">
      <div className="flex flex-col items-center gap-6 bg-slate-800 border border-slate-700 rounded-xl p-4 max-w-2xl w-full shadow-2xl">
        {/* Icon */}
        <div className="p-6 bg-slate-700 rounded-2xl border border-slate-600">
          <MdMiscellaneousServices className="w-20 h-20 text-sky-500" />
        </div>

        {/* Title */}
        <h1 className="text-4xl font-extrabold text-white text-center">
          Services
        </h1>

        {/* Coming Soon Badge */}
        <div className="flex items-center gap-2 bg-sky-600/20 border border-sky-500/50 rounded-lg px-4 py-2">
          <FaClock className="w-5 h-5 text-sky-400" />
          <span className="text-sky-400 font-semibold">Coming Soon</span>
        </div>

        {/* Description */}
        <p className="text-slate-400 text-center text-lg leading-relaxed">
          The Services page is currently under development. This feature will allow you to manage and monitor your services that you configured in server settings.
        </p>

          {/* Features List */}
          <div className="w-full bg-slate-700/50 border border-slate-600 rounded-lg p-6 mt-4">
            <div className="flex items-center gap-2 mb-4">
              <FaTools className="w-5 h-5 text-sky-500" />
              <h3 className="text-lg font-semibold text-white">Planned Features</h3>
            </div>
            <ul className="space-y-3 text-slate-300">
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>Start, stop, and restart services</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>Monitor service status and health</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>Configure service settings and parameters</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>View service logs and diagnostics</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>Run commands inside service</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>Permission-based access control</span>
              </li>
            </ul>
          </div>

          {/* Use case List */}
          <div className="w-full bg-slate-700/50 border border-slate-600 rounded-lg p-6 mt-4">
            <div className="flex items-center gap-2 mb-4">
              <BiSolidParty className="w-5 h-5 text-sky-500" />
              <h3 className="text-lg font-semibold text-white">Use Cases</h3>
            </div>
            <ul className="space-y-3 text-slate-300">
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>Web servers</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>Game servers</span>
              </li>
              <li className="flex items-start gap-2">
                <span className="text-sky-500 mt-1">•</span>
                <span>And much more...</span>
              </li>
            </ul>
          </div>

        {/* Footer Message */}
        <p className="text-slate-500 text-sm text-center mt-4">
          Stay tuned for updates! This feature will be available in a future release.
        </p>
      </div>
    </div>
  )
}

export default Services