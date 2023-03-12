import { mount } from "@vue/test-utils";
import ViewPage from "@/components/Pages/View.vue";
import { createStore } from "vuex";

describe("ViewPage.vue", () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });

  const mockRoute = {
    params: {
      id: 1,
    },
  };

  const createVuexStore = () => {
    return createStore({
        state() {
            return {
              Pages: {
                onePage: {
                  ID: "1",
                  title: "Post 1",
                  body: "A simple page text body",
                  slug: "simple-slug-text",
                },
              },
            };
          },
      mutations: {
        updatePage(state, res) {
            state.Pages.onePage = res
        },
      },
      actions: {
        getPage(ctx, res) {
          ctx.commit("updatePage", res);
        },
      },
    });
  };

  function factory() {
    const store = createVuexStore();

    return mount(ViewPage, {
      global: {
        plugins: [store],
        mocks: {
           $route: mockRoute,
        },
      },
    });
  }

  it("renders props.msg when passed", async () => {
    const msg =
      "AÇÕES Post 1 Visualização de página do siteA simple page text bodyEditarRemover";
    const wrapper = factory();

    console.log(wrapper.emitted())

    expect(mockRoute.params.id).toBe(1);
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
