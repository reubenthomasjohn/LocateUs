import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom"; // Import useNavigate

function SignInForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const submit = async () => {
    const api = "https://locatetogether.net/members/users/login"; // Update with your actual API
    try {
      const response = await axios.post(api, {
        username: username, // Use state values
        password: password,
      });
      if (response.status === 200 && response.data.access_token) {
        console.log("Login successful:", response.data);
        localStorage.setItem("access_token", response.data.access_token); // Store token for authentication
        navigate("/dashboard"); // Redirect to dashboard
      }
    } catch (err) {
      console.error("User login failed:", err);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen">
      <Card className="mx-auto max-w-sm">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl font-bold">Login</CardTitle>
          <CardDescription>
            Please enter your username and password.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="username">Username</Label>
            <Input
              id="username"
              type="text"
              placeholder="Username"
              required
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="password">Password</Label>
            <Input
              id="password"
              type="password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <Button type="submit" className="w-full" onClick={submit}>
            Log In
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}

export default SignInForm;
