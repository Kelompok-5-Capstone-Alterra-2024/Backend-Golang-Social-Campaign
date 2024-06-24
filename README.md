<div id="top"></div>

<!-- PROJECT LOGO -->
<br/>
<div align="center">
  <!--  Link to the repository -->
  <a href="https://github.com/Kelompok-5-Capstone-Alterra-2024">
    <img src="https://github.com/Kelompok-5-Capstone-Alterra-2024/Backend-Golang-Social-Campaign/assets/114483889/f521c419-ab65-411e-a02c-97df20c89f29" width="70" height="70
    ">
  </a>

  <h3 align="center">Peduli Pintar</h3>

  <p align="center">
    Project Capstone Program Immersive Alterra Academy
    <br />
    <a href="[https://documenter.getpostman.com/view/21327885/2sA3QpBswH](https://documenter.getpostman.com/view/21327885/2sA3QpBswH#intro)"><strong>Explore the docs Open API »</strong></a>
    <br />
  </p>
</div>


<!-- ABOUT THE PROJECT -->
## About The Project
**Peduli Lindungi** is a platform designed to facilitate the community in donating, volunteering, and reading articles related to social and humanitarian activities. The platform aims to build a caring and active community involved in various social activities.

### Features Overview
In the application, there are 2 roles: User and Admin.

</br>

**User Endpoints**
| Feature                        | Endpoint Description                          |
|--------------------------------|-----------------------------------------------|
| Login                          | Login                                         |
| Register                       | Register                                      |
| Forget Password                | Forget Password                               |
| Reset Password with OTP        | Reset Password with OTP                       |
| Profile                        | Get User Profile                              |
|                                | Edit Profile                                  |
|                                | Change Password                               |
|                                | Bookmark Fundraising                          |
|                                | Bookmark Article                              |
|                                | Bookmark Volunteer                            |
| Fundraising                    | Get All Category                              |
|                                | Get Fundraising by Category                   |
|                                | Get All Fundraising                           |
|                                | Get Top Fundraisings                          |
|                                | Donation (Payment Gateway)                    |
|                                | Post Donation (Manual)                        |
|                                | Post Like Comment Donation                    |
|                                | Delete Unlike Comment Donation                |
|                                | Get Detail Fundraising                        |
|                                | Get Detail Organization                       |
| Volunteer                      | Register Volunteer                            |
|                                | Confirm Volunteer Apply                       |
|                                | Apply For Volunteer                           |
|                                | Get Detail Volunteer                          |
|                                | Get All Volunteer                             |
|                                | Get Top Volunteer                             |
|                                | Get Detail Organization                       |
| Article                        | Get All Article                               |
|                                | Get Article by Id                             |
|                                | Get Top Article                               |
|                                | Create Comment                                |
|                                | Get All Comment By Article ID                 |
|                                | Like Comment & Unlike Comment                 |
| History                        | Get All Volunteer                             |
|                                | Get Detail Volunteer                          |
|                                | Get All Fundraising                           |
|                                | Get Detail Fundraising                        |

</br>


**Admin Endpoints**
| Feature                        | Endpoint Description                          |
|--------------------------------|-----------------------------------------------|
| Login                          | Login                                         |
| Users                          | Get All Users                                 |
|                                | Get User Donations                            |
|                                | Delete User                                   |
| Fundraising                    | Get All                                       |
|                                | Create New Fundraising                        |
|                                | Get Detail Fundraising                        |
|                                | Get Donations by Fundraising Id               |
|                                | Delete Fundraising By Id                      |
|                                | Edit Fundraising                              |
| Organization                   | Add Organization                              |
|                                | Get All Organizations                         |
|                                | Edit Organization                             |
|                                | Delete Organization                           |
| Article                        | Get All Article                               |
|                                | Post Article                                  |
|                                | Update Article                                |
|                                | Delete Article                                |
|                                | Get Detail                                    |
|                                | Get All Comment By Article ID                 |
| Volunteer                      | Create Volunteer Vacancy                      |
|                                | Get All Volunteer Vacancy                     |
|                                | Update Volunteer Vacancy                      |
|                                | Get Detail                                    |
|                                | Get All Apply by Volunteer ID                 |
| Donation                       | Get All History                               |
|                                | Add Amount to Donation                        |

### Built With
![VS Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=FFFFFF)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=FFFFFF)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=FFFFFF)&nbsp;
![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman&logoColor=FFFFFF)&nbsp;

<!-- How to Use -->
## How to Use
**Setup on Local Server**
- Install Golang, Postman, MySQL Workbench
- Clone repository with HTTPS:
```
git clone https://github.com/Kelompok-5-Capstone-Alterra-2024/Backend-Golang-Social-Campaign.git
```
* Create File `.env`:
```
DB_USER=root
DB_PASSWORD=@Kris0624Capstone5
DB_HOST=34.126.74.53
DB_PORT=3306
DB_NAME=capstone5
```
* Run `main.go` on local terminal
```
$ go run main.go
```
* Run the endpoint according to the [OpenAPI Documentation](https://documenter.getpostman.com/view/21327885/2sA3QpBswH) via Postman 

<!-- ERD -->
## ERD
<a href="https://lucid.app/lucidchart/fe4c6c45-5053-467f-8fec-b1c0d1d9c291/edit?viewport_loc=-1805%2C-1184%2C7742%2C3878%2CSfR4hQLHwG25&invitationId=inv_5f45b9e5-78db-4caa-95d2-366a5308b19c">
    <img src="https://github.com/Kelompok-5-Capstone-Alterra-2024/Backend-Golang-Social-Campaign/assets/114483889/a4827381-1c0e-498d-8658-5ec6fe59d327" width="900" height="900">
</a>

<!-- end -->
