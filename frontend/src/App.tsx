import { useEffect, useState } from "react";
import { Map } from "./components/Map";
import { UserTable } from "./components/UserTable";
import { Search } from "lucide-react";
import { Header } from "./components/Header";
import { Footer } from "./components/Footer";
import { ButtonWithIcon } from "./components/ui/IconButton";
import { Analytics, IAnalytics } from "./components/Analytics";
import { User } from "./store/atoms/users";
import { useRecoilState } from "recoil";
import { userState } from "./store/atoms/users";
import axios from "axios";

function App() {
  const [users, setUsers] = useRecoilState(userState);

  const initialAnalytics: IAnalytics = {
    messagesSent: "0",
    responsesReceived: "0",
    averageResponseTime: "0s",
    responseRate: "0",
  };
  const [analytics, setAnalytics] = useState<IAnalytics>(initialAnalytics);
  const [searchQuery, setSearchQuery] = useState("");

  useEffect(() => {
    axios
      .get("http://localhost:8080/users")
      .then((response) => setUsers(response.data))
      .catch((err) => console.error("User fetch failed:", err));

    //TODO:: // const analytics = axios.get("http://locahost:8080/analytics");

    // Sample analytics data
    const analytics = {
      messagesSent: "1250",
      responsesReceived: "1180",
      averageResponseTime: "1.5",
      responseRate: "94.4",
    };
    setAnalytics(analytics);
  }, []);

  const handleDeleteUser = (userId: number) => {
    setUsers(users.filter((user: User) => user.id !== userId));
  };

  const handleEditUser = (userId: number) => {
    // In a real application, this would open a modal or navigate to an edit page
    console.log("Edit user:", userId);
  };

  const filteredUsers = users.filter((user) => {
    const searchLower = searchQuery.toLowerCase();
    return (
      user.full_name.toLowerCase().includes(searchLower) ||
      user.phone_number.includes(searchQuery)
    );
  });

  return (
    <div className="min-h-screen flex flex-col bg-gray-50">
      <Header />
      <main className="flex-grow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div className="space-y-8">
            <div className="flex flex-row justify-between">
              <div>
                <h1 className="text-3xl font-bold text-gray-900">
                  Location Dashboard
                </h1>
                <p className="mt-2 text-sm text-gray-500">
                  Track member locations and manage member data
                </p>
              </div>
              <div className="m-4">
                <ButtonWithIcon>Refresh Database</ButtonWithIcon>
              </div>
            </div>

            <Analytics analytics={analytics} users={users} />

            <div className="bg-white shadow rounded-lg p-6">
              <h2 className="text-lg font-medium text-gray-900 mb-4">
                Location Heatmap
              </h2>
              <Map users={users} />
            </div>

            <div className="bg-white shadow rounded-lg">
              <div className="px-6 py-4 border-b border-gray-200">
                <div className="flex justify-between items-center">
                  <h2 className="text-lg font-medium text-gray-900">Members</h2>
                  <div className="relative">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <Search className="h-5 w-5 text-gray-400" />
                    </div>
                    <input
                      type="text"
                      placeholder="Search members..."
                      value={searchQuery}
                      onChange={(e) => setSearchQuery(e.target.value)}
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                  </div>
                </div>
              </div>
              <UserTable
                users={filteredUsers}
                onEdit={handleEditUser}
                onDelete={handleDeleteUser}
              />
            </div>
          </div>
        </div>
      </main>
      <Footer />
    </div>
  );
}

export default App;
