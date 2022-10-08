import { ViewcontactComponent } from './component/viewcontact/viewcontact.component';
import { RegisterComponent } from './component/register/register.component';
import { HomeComponent } from './component/home/home.component';
import { LoginComponent } from './component/login/login.component';
import { NgModule, Component } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LogoutComponent } from './component/logout/logout.component';
import { AdressbookComponent } from './adressbook/adressbook.component';
import { AddcontactComponent } from './component/addcontact/addcontact.component';

const routes: Routes = [
  // { path: '', redirectTo: '/', component: HomeComponent },
  { path: '', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'logout', component: LogoutComponent },
  { path: 'address', component: AdressbookComponent },
  { path: 'add', component: AddcontactComponent },
  { path: 'view', component: ViewcontactComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
