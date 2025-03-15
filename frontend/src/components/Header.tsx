export const Header = () => {
  return (
    <header className="bg-white shadow-sm">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-4">
            <div className="flex items-center justify-center w-40 h-10 rounded-lg">
              <img src="https://media.christcommunitychurch.in/sites/2/2017/06/2017-logo-2.png"></img>
              {/* <Building2 className="w-6 h-6 text-white" /> */}
            </div>
            {/* <span className="text-xl font-bold text-gray-900">
              Christ Community Church
            </span> */}
          </div>
          <nav className="flex space-x-4">
            <a href="#" className="text-gray-500 hover:text-gray-900">
              Dashboard
            </a>
            <a href="#" className="text-gray-500 hover:text-gray-900">
              Settings
            </a>
            <a href="#" className="text-gray-500 hover:text-gray-900">
              Help
            </a>
            <a
              onClick={() => {
                localStorage.removeItem("access_token");
              }}
              href="/"
              className="text-gray-500 hover:text-gray-900"
            >
              Logout
            </a>
          </nav>
        </div>
      </div>
    </header>
  );
};
