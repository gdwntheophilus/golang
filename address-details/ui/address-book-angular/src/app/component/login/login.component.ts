import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UrlsService } from '../../url/urls.service';
import { User } from '../../user/user';
import { FormControl, Validators } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  // ngOnInit(): void {}
  OnInit(): void {}
  urlsService: any;
  user: User = {
    emailid: '',
    password: '',
  };
  loginStatus = '';

  constructor(private httpClient: HttpClient, private router: Router) {
    this.urlsService = new UrlsService();
  }
  ngOnInit(): void {
    throw new Error('Method not implemented.');
  }

  UserLoginSubmit() {
    // console.log(this.user.emailid);
    // console.log(this.user.password);
    // console.log(this.urlsService.urls().getRegisteredUser);
    var headers = new HttpHeaders({
      'Content-Type': 'application/json',
    });

    if (this.user.emailid != '' && this.user.password != '') {
      console.log(this.user);
      this.httpClient
        .post(this.urlsService.urls().getRegisteredUser, this.user, { headers })
        .subscribe((response: any) => {
          // console.log(response);
          // console.log(response.status);
          if (response.status === 'error') {
            this.loginStatus = response.message;
          }
          if (response.status === 'Ok') {
            this.loginStatus = response.message;
            console.log('login router');
            sessionStorage.setItem('token', this.user.emailid);
            this.router.navigate(['address']);
          }
        });
    } else {
      console.log(this.user);
      this.loginStatus = 'Username or Password is null!!';
      console.log(this.loginStatus);
    }
  }
}
