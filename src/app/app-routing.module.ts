import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ApiReferenceComponent } from './api-reference/api-reference.component';
import { AppComponent } from './app.component';
import { LoadDocComponent } from './load-doc/load-doc.component';

const routes: Routes = [
  { path: '', component: AppComponent, pathMatch: 'full' },
  {
    path: 'api-reference',
    component: ApiReferenceComponent,
    children: [{ path: '**', component: ApiReferenceComponent }],
  },
  {
    path: 'load-doc',
    component: LoadDocComponent,
    children: [{ path: '**', component: LoadDocComponent }],
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
