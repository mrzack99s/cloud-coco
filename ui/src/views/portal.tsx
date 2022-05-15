import Navbar from "../components/navbar";
import "../assets/custom.css"
import AdminOverview from "../components/portals/administrator/overview";
import UserOverview from "../components/portals/users/overview";
import { useCookies } from 'react-cookie';

export default () => {
	const [cookies, setCookie, removeCookie] = useCookies(['role']);

	return (
		<div>
			<Navbar />
			<div className="grid m-0" style={{ position: "relative", top: "65px" }}>
				<div className="col-fixed" style={{ width: "10%" }}></div>
				<div className="col sm:12" style={{}}>
					{cookies.role == "administrator" &&
						<AdminOverview />
					}
					{cookies.role != "administrator" &&
						<UserOverview />
					}
				</div>
				<div className="col-fixed" style={{ width: "10%" }}></div>
			</div>
		</div>
	);
};
